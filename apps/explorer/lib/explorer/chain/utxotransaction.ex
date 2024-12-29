defmodule Explorer.Chain.UTXOTransaction do
  @moduledoc "Models a Web3 transaction."

  use Explorer.Schema

  require Logger
  alias Explorer.Repo
  import Ecto.Query, only: [from: 2, preload: 3, subquery: 1, where: 3]
  import Explorer.Chain.UTXOAddress, only: [utxoaddress_update: 2]
  alias Ecto.Changeset

  alias Explorer.Chain.{
    Block,
    Hash,
    Transaction
  }

  @optional_attrs ~w(block_hash blockorder txindex locktime spenttxhash txtime status fee)a

  @required_attrs ~w(hash index size toaddress amount vin pkscript)a

  @typedoc """
  The index of the transaction in its block.
  """
  @type transaction_index :: non_neg_integer()

  @typedoc """
  `t:standard_v/0` + `27`

  | `v`  | X      | Y    |
  |------|--------|------|
  | `27` | lower  | even |
  | `28` | lower  | odd  |
  | `29` | higher | even |
  | `30` | higher | odd  |

  **Note: that `29` and `30` are exceedingly rarely, and will in practice only ever be seen in specifically generated
  examples.**
  """
  @type v :: 27..30

  @typedoc """
  How much the sender is willing to pay in wei per unit of gas.
  """
  @type wei_per_gas :: Wei.t()

  @type t :: %__MODULE__{
          block_hash: String.t() | nil,
          blockorder: Block.blockorder() | nil,
          size: non_neg_integer(),
          txindex: non_neg_integer(),
          index: non_neg_integer(),
          hash: String.t(),
          locktime: non_neg_integer() | nil,
          toaddress: String.t(),
          amount: Decimal.t(),
          fee: Decimal.t(),
          spenttxhash: String.t() | nil,
          txtime: DateTime.t(),
          vin: String.t(),
          pkscript: String.t(),
          status: non_neg_integer()
        }

  @primary_key false
  @unique_index [:hash, :index]
  schema "utxotransactions" do
    field(:block_hash, :string)
    field(:blockorder, :integer)
    field(:size, :integer)
    field(:txindex, :integer)
    field(:index, :integer)
    field(:hash, :string)
    field(:locktime, :integer)
    field(:toaddress, :string)
    field(:amount, :decimal)
    field(:fee, :decimal)
    field(:spenttxhash, :string)
    field(:txtime, :utc_datetime_usec)
    field(:vin, :string)
    field(:pkscript, :string)
    field(:status, :integer)

    timestamps()
  end

  def changeset(%__MODULE__{} = transaction, attrs \\ %{}) do
    attrs_to_cast = @required_attrs ++ @optional_attrs

    transaction
    |> cast(attrs, attrs_to_cast)
    |> validate_required(@required_attrs)
  end

  def not_pending_transactions(query) do
    where(query, [t], not is_nil(t.blockorder))
  end
  def insert_tx(%{"hash": hash} = attrs) do
    utxoaddress_update(attrs.toaddress, attrs.amount)
    %__MODULE__{}
    |> changeset(attrs)
    |> Repo.insert!()

  end

  def insert_tx(%{:error => err}) do
    {:error, err}
  end

  @error_message "can't be set when status is not :error"

  @doc """
  Builds an `Ecto.Query` to fetch transactions with the specified blockorder
  """
  def transactions_with_blockorder(blockorder) do
    from(
      t in Transaction,
      where: t.blockorder == ^blockorder
    )
  end

  def utxotx_update_status(hash, index, spenttxhash) do
    tx = Repo.one(from u in __MODULE__, where: u.hash == ^hash and u.index == ^index)
    IO.inspect(tx)
    case tx do
      nil ->
        {:error, "tx not found"}

      tx ->
        changeset = Ecto.Changeset.change(tx, spenttxhash: spenttxhash, status: 1)

        case Repo.update(changeset) do
          {:ok, updated_tx} ->
            utxoaddress_update(tx.toaddress, -tx.amount)
            {:ok, updated_tx}

          {:error, changeset} ->
            {:error, changeset.errors}
        end
    end
  end
end
