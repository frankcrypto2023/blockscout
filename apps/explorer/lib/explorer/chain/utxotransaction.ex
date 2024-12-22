defmodule Explorer.Chain.UTXOTransaction do
  @moduledoc "Models a Web3 transaction."

  use Explorer.Schema

  require Logger

  import Ecto.Query, only: [from: 2, preload: 3, subquery: 1, where: 3]

  alias Ecto.Changeset

  alias Explorer.Chain.{
    Block,
    Hash,
    Transaction
  }

  @optional_attrs ~w(block_hash block_number status)a

  @required_attrs ~w(hash)a

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
          block_hash: Hash.t() | nil,
          block_number: Block.block_number() | nil,
          size: non_neg_integer(),
          hash: Hash.t(),
          status: non_neg_integer(),
          locketime: NaiveDateTime.t() | nil,
        }

  @primary_key {:hash, Hash.Full, autogenerate: false}
  schema "utxotransactions" do
    field(:block_number, :integer)
    field(:size, :integer)
    field(:index, :integer)
    field(:status, :integer)
    field(:locketime, :integer)
    field(:block_hash, Hash.Full)

    timestamps()
  end

  def changeset(%__MODULE__{} = transaction, attrs \\ %{}) do
    attrs_to_cast = @required_attrs ++ @optional_attrs

    transaction
    |> cast(attrs, attrs_to_cast)
    |> validate_required(@required_attrs)
    |> validate_error()
    |> validate_status()
    |> check_error()
    |> check_status()
    |> foreign_key_constraint(:block_hash)
    |> unique_constraint(:hash)
  end

  def not_pending_transactions(query) do
    where(query, [t], not is_nil(t.block_number))
  end

  @error_message "can't be set when status is not :error"

  defp check_error(%Changeset{} = changeset) do
    check_constraint(changeset, :error, message: @error_message, name: :error)
  end

  @status_message "can't be set when the block_hash is unknown"

  defp check_status(%Changeset{} = changeset) do
    check_constraint(changeset, :status, message: @status_message, name: :status)
  end

  defp validate_error(%Changeset{} = changeset) do
    if Changeset.get_field(changeset, :status) != :error and Changeset.get_field(changeset, :error) != nil do
      Changeset.add_error(changeset, :error, @error_message)
    else
      changeset
    end
  end

  defp validate_status(%Changeset{} = changeset) do
    if Changeset.get_field(changeset, :block_hash) == nil and
         Changeset.get_field(changeset, :status) != nil do
      Changeset.add_error(changeset, :status, @status_message)
    else
      changeset
    end
  end

  @doc """
  Builds an `Ecto.Query` to fetch transactions with the specified block_number
  """
  def transactions_with_block_number(block_number) do
    from(
      t in Transaction,
      where: t.block_number == ^block_number
    )
  end

end
