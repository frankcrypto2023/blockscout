defmodule Explorer.Repo.Migrations.CreateUTXOTransactions do
  use Ecto.Migration

  def change do
    create table(:utxotransactions, primary_key: false) do
      add(:block_hash, references(:utxoblocks, column: :hash, on_delete: :delete_all, type: :bytea), null: true)
      add(:blockorder, :integer, null: true)
      # `null` when a pending transaction
      add(:size, :integer,  null: true)
      add(:txindex, :integer,  null: true)
      add(:index, :integer,  null: true)
      add(:hash, :bytea, null: false)

      add(:locktime, :integer, null: true)
      add(:toaddress, :bytea, null: false)
      add(:amount, :numeric, precision: 100, null: false)
      add(:fee, :numeric, precision: 100, null: true)
      add(:txtime, :utc_datetime_usec, null: true)
      add(:vin,:bytea, null: false)
      add(:pkscript, :bytea, null: false)
      add(:spenttxhash,:bytea, null: true)

      # `null` when a pending transaction
      add(:status, :integer, null: true)


      timestamps(null: false, type: :utc_datetime_usec)

      # `null` when a pending transaction

      # `null` when a pending transaction
      # denormalized from `blocks.number` to improve `Explorer.Chain.recent_collated_transactions/0` performance

    end

    create(unique_index(:utxotransactions, [:hash,:index]))
  end
end
