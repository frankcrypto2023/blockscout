defmodule Explorer.Repo.Migrations.CreateUTXOBlocks do
  use Ecto.Migration

  def change do
    create table(:utxoblocks, primary_key: false) do
      add(:txsvalid, :boolean, null: false)
      add(:difficulty, :numeric, precision: 50)
      add(:hash, :bytea, null: false, primary_key: true)
      add(:miner_hash, :bytea, null: false)
      add(:nonce, :bytea, null: false)
      add(:height, :bigint, null: false)
      add(:weight, :bigint, null: false)
      add(:order, :bigint, null: false)
      add(:txns, :bigint, null: false)
      add(:powname, :bytea, null: false)
      add(:status, :numeric) # 0 = pending, 1 = valid, 2 = invalid
      # not a foreign key to allow skipped blocks
      add(:parent_root, :bytea, null: false)
      add(:tx_root, :bytea, null: false)
      add(:parents, :bytea, null: false)
      add(:children, :bytea, null: false)

      add(:timestamp, :utc_datetime_usec, null: false)

      timestamps(null: false, type: :utc_datetime_usec)
    end

    create(index(:utxoblocks, [:timestamp]))
    create(index(:utxoblocks, [:hash], unique: true, where: ~s(txsvalid), name: :one_txsvalid_child_per_hash))
    create(index(:utxoblocks, [:order], unique: true, where: ~s(txsvalid), name: :one_txsvalid_block_at_order))
  end
end
