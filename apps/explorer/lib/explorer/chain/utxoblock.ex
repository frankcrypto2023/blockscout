defmodule Explorer.Chain.UTXOBlock do
  @moduledoc """
  A package of data that contains zero or more transactions, the hash of the previous block ("parent"), and optionally
  other data. Because each block (except for the initial "genesis block") points to the previous block, the data
  structure that they form is called a "blockchain".
  """

  use Explorer.Schema

  alias Explorer.Chain.{ Hash, }

  @optional_attrs ~w(difficulty)a

  @required_attrs ~w(txsvalid hash miner_hash nonce order height parent_root timestamp)a

  @typedoc """
  How much work is required to find a hash with some number of leading 0s.  It is measured in hashes for PoW
  (Proof-of-Work) chains like Ethereum.  In PoA (Proof-of-Authority) chains, it does not apply as blocks are validated
  in a round-robin fashion, and so the value is always `Decimal.new(0)`.
  """
  @type difficulty :: Decimal.t()

  @typedoc """
  Number of the block in the chain.
  """
  @type block_number :: non_neg_integer()

  @typedoc """
   * `txsvalid`
     * `true` - this is a block which txs are valid
     * `false` - this is block which txs are invalid, maybe red block
   * `difficulty` - how hard the block was to mine.
   * `hash` - the hash of the block.
   * `miner_hash` - the base58 of the address.
   * `nonce` - the hash of the generated proof-of-work.  Not used in Proof-of-Authority chains.
   * `number` - which block this is along the chain.
   * `parent_root` - the hash of the parent block, which should have the previous `number`
   * `powname` - the name of the pow
   * `timestamp` - When the block was collated
   * `txns` - the transactions count in this block.
  """
  @type t :: %__MODULE__{
          txsvalid: boolean(),
          difficulty: difficulty(),
          hash: Hash.Full.t(),
          miner_hash: String.t(),
          nonce: non_neg_integer(),
          order: block_number(),
          height: block_number(),
          parent_root: Hash.t(),
          timestamp: DateTime.t(),
          powname: String.t(),
          txns: non_neg_integer(),
        }

  @primary_key {:hash, Hash.Full, autogenerate: false}
  schema "utxoblocks" do
    field(:difficulty, :decimal)
    field(:miner_hash, :string)
    field(:nonce, Hash.Nonce)
    field(:order, :integer)
    field(:height, :integer)
    field(:timestamp, :utc_datetime_usec)
    field(:parent_root, Hash.Full)
    field(:powname, :string)
    field(:txns, :integer)
    field(:txsvalid, :boolean)

    timestamps()
  end

  def changeset(%__MODULE__{} = block, attrs) do
    block
    |> cast(attrs, @required_attrs ++ @optional_attrs)
    |> validate_required(@required_attrs)
    |> unique_constraint(:hash, name: :utxoblocks_pkey)
  end

  def number_only_changeset(%__MODULE__{} = block, attrs) do
    block
    |> cast(attrs, @required_attrs ++ @optional_attrs)
    |> validate_required([:number])
    |> unique_constraint(:hash, name: :utxoblocks_pkey)
  end

end
