defmodule EthereumJSONRPC.UTXOBlock do
  @moduledoc """
  Block format as returned by [`eth_getBlockByHash`](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash)
  and [`eth_getBlockByNumber`](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber).
  """

  import EthereumJSONRPC, only: [quantity_to_integer: 1, timestamp_to_datetime: 1]

  alias EthereumJSONRPC.{UTXOTransaction,UTXOTransactionOutput}

  @type elixir :: %{String.t() => non_neg_integer | DateTime.t() | String.t() | nil}
  @type params :: %{
          txsvalid: boolean(),
          difficulty: pos_integer(),
          hash: EthereumJSONRPC.hash(),
          miner_hash: String.t(),
          order: non_neg_integer(),
          height: non_neg_integer(),
          parentroot: EthereumJSONRPC.hash(),
          timestamp: DateTime.t(),
          pow: %{
            pow_name: String.t(),
            pow_type: non_neg_integer(),
            nonce: non_neg_integer(),
          },
          transactions: [%{
            txid: EthereumJSONRPC.hash(),
            size: non_neg_integer(),
            locktime: non_neg_integer(),
            txsvalid: boolean(),
            vout: [
              %{
                amount: non_neg_integer(),
                scriptPubKey: %{
                  addresses: [String.t()],
                }
              }
            ]
          }],
        }

  @type t :: %{String.t() => EthereumJSONRPC.data() | EthereumJSONRPC.hash() | EthereumJSONRPC.quantity() | nil}

  def from_response(%{id: id, result: nil}, id_to_params) when is_map(id_to_params) do
    params = Map.fetch!(id_to_params, id)

    {:error, %{code: 404, message: "Not Found", data: params}}
  end

  def from_response(%{id: id, result: block}, id_to_params) when is_map(id_to_params) do
    true = Map.has_key?(id_to_params, id)

    {:ok, block}
  end

  def from_response(%{id: id, error: error}, id_to_params) when is_map(id_to_params) do
    params = Map.fetch!(id_to_params, id)
    annotated_error = Map.put(error, :data, params)

    {:error, annotated_error}
  end

  @spec elixir_to_params(elixir) :: params
  def elixir_to_params(
        %{
          "txsvalid": txsvalid,
          "difficulty": difficulty,
          "hash": hash,
          "order": order,
          "height": height,
          "parentroot": parentroot,
          "timestamp": timestamp,
          "pow": pow,
          "transactions": transactions
        } = elixir
      ) do
    %{
      txsvalid: txsvalid,
      difficulty: difficulty,
      hash: hash,
      miner_hash: String.t(),
      nonce: Map.get(pow, "nonce", 0),
      order: order,
      height: height,
      parentroot: parentroot,
      timestamp: timestamp,
      pow: pow,
      transactions: transactions
    }
  end

  @spec elixir_to_transactions(elixir) :: Transactions.elixir()
  def elixir_to_transactions(%{"transactions" => transactions}), do: transactions

  def elixir_to_transactions(_), do: []

  def to_elixir(block) when is_map(block) do
    block
  end
end
