defmodule BlockScoutWeb.API.UTXO.AddressView do
  use BlockScoutWeb, :view

  import BlockScoutWeb.Account.AuthController, only: [current_user: 1]

  alias BlockScoutWeb.AddressView
  alias BlockScoutWeb.API.UTXO.{ApiView, Helper}
  alias BlockScoutWeb.API.UTXO.Helper
  alias Explorer.{Chain, Market}
  alias Explorer.Chain.Address.Counters
  alias Explorer.Chain.{Address, SmartContract}

  @api_true [api?: true]

  def render("message.json", assigns) do
    ApiView.render("message.json", assigns)
  end

  def render("address.json", %{address: address, conn: conn}) do
    prepare_address(address, conn)
  end

  def render("coin_balance.json", %{coin_balance: coin_balance}) do
    prepare_coin_balance_history_entry(coin_balance)
  end

  def render("coin_balances.json", %{coin_balances: coin_balances, next_page_params: next_page_params}) do
    %{"items" => Enum.map(coin_balances, &prepare_coin_balance_history_entry/1), "next_page_params" => next_page_params}
  end

  def render("coin_balances_by_day.json", %{coin_balances_by_day: coin_balances_by_day}) do
    Enum.map(coin_balances_by_day, &prepare_coin_balance_history_by_day_entry/1)
  end

  def render("addresses.json", %{
        addresses: addresses,
        next_page_params: next_page_params,
        exchange_rate: exchange_rate,
        total_supply: total_supply
      }) do
    %{
      items: Enum.map(addresses, &prepare_address/1),
      next_page_params: next_page_params,
      exchange_rate: exchange_rate.usd_value,
      total_supply: total_supply && to_string(total_supply)
    }
  end

  def prepare_address({address, nonce}) do
    nil
    |> Helper.address_with_info(address, address.hash, true)
    |> Map.put(:tx_count, to_string(nonce))
    |> Map.put(:coin_balance, if(address.fetched_coin_balance, do: address.fetched_coin_balance.value))
  end

  def prepare_address(address, conn \\ nil) do
    base_info = Helper.address_with_info(conn, address, address.hash, true)
    is_proxy = AddressView.smart_contract_is_proxy?(address, @api_true)

    {implementation_address, implementation_name} =
      with true <- is_proxy,
           {address, name} <- SmartContract.get_implementation_address_hash(address.smart_contract, @api_true),
           false <- is_nil(address),
           {:ok, address_hash} <- Chain.string_to_address_hash(address),
           checksummed_address <- Address.checksum(address_hash) do
        {checksummed_address, name}
      else
        _ ->
          {nil, nil}
      end

    balance = address.fetched_coin_balance && address.fetched_coin_balance.value
    exchange_rate = Market.get_coin_exchange_rate().usd_value

    creator_hash = AddressView.from_address_hash(address)
    creation_tx = creator_hash && AddressView.transaction_hash(address)

    write_custom_abi? = AddressView.has_address_custom_abi_with_write_functions?(conn, address.hash)
    read_custom_abi? = AddressView.has_address_custom_abi_with_read_functions?(conn, address.hash)

    Map.merge(base_info, %{
      "creator_address_hash" => creator_hash && Address.checksum(creator_hash),
      "creation_tx_hash" => creation_tx,
      "coin_balance" => balance,
      "exchange_rate" => exchange_rate,
      "implementation_name" => implementation_name,
      "implementation_address" => implementation_address,
      "block_number_balance_updated_at" => address.fetched_coin_balance_block_number,
      "has_custom_methods_read" => read_custom_abi?,
      "has_custom_methods_write" => write_custom_abi?,
      "has_methods_read" => AddressView.smart_contract_with_read_only_functions?(address),
      "has_methods_write" => AddressView.smart_contract_with_write_functions?(address),
      "has_methods_read_proxy" => is_proxy,
      "has_methods_write_proxy" => AddressView.smart_contract_with_write_functions?(address) && is_proxy,
      "has_decompiled_code" => AddressView.has_decompiled_code?(address),
      "has_validated_blocks" => Counters.check_if_validated_blocks_at_address(address.hash, @api_true),
      "has_logs" => Counters.check_if_logs_at_address(address.hash, @api_true),
      "has_tokens" => Counters.check_if_tokens_at_address(address.hash, @api_true),
      "has_token_transfers" => Counters.check_if_token_transfers_at_address(address.hash, @api_true),
      "watchlist_address_id" => Chain.select_watchlist_address_id(get_watchlist_id(conn), address.hash),
      "has_beacon_chain_withdrawals" => Counters.check_if_withdrawals_at_address(address.hash, @api_true)
    })
  end

  def prepare_coin_balance_history_entry(coin_balance) do
    %{
      "transaction_hash" => coin_balance.transaction_hash,
      "block_number" => coin_balance.block_number,
      "delta" => coin_balance.delta,
      "value" => coin_balance.value,
      "block_timestamp" => coin_balance.block_timestamp
    }
  end

  def prepare_coin_balance_history_by_day_entry(coin_balance_by_day) do
    %{
      "date" => coin_balance_by_day.date,
      "value" => coin_balance_by_day.value
    }
  end

  def get_watchlist_id(conn) do
    case current_user(conn) do
      %{watchlist_id: wl_id} ->
        wl_id

      _ ->
        nil
    end
  end
end
