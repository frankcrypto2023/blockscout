defmodule Explorer.Repo.Migrations.CreateUTXOTransactions do
  use Ecto.Migration

  def change do
    create table(:utxotransactions, primary_key: false) do
      # `null` when a pending transaction
      add(:size, :numeric, precision: 100, null: true)
      add(:txvalid, :boolean, null: false)
      # `null` before internal transactions are fetched or if no error in those internal transactions
      add(:error, :string, null: true)

      add(:locktime, :bigint, null: false)

      # txid
      add(:hash, :bytea, null: false, primary_key: true)

      add(:type, :string, null: true)

      # `null` when a pending transaction
      add(:status, :integer, null: true)


      timestamps(null: false, type: :utc_datetime_usec)

      # `null` when a pending transaction
      add(:block_hash, references(:utxoblocks, column: :hash, on_delete: :delete_all, type: :bytea), null: true)

      # `null` when a pending transaction
      # denormalized from `blocks.number` to improve `Explorer.Chain.recent_collated_transactions/0` performance
      add(:block_number, :integer, null: true)

    end

    create(unique_index(:utxotransactions, [:hash]))
  end
end
