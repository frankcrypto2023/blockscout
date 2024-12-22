defmodule Explorer.Repo.Migrations.CreateUTXOAddresstxs do
  use Ecto.Migration

  def change do
    create table(:utxoaddresstxs, primary_key: false) do
      add(:address, :string, null: true)
      add(:vout, :integer, null: false)
      add(:value, :bigint,  null: false)
      add(:pkscript, :bytea,  null: false)
      add(:vins, :string, null: false)
      add(:status, :integer, null: true)
      add(:txtime, :utc_datetime_usec, null: false)
      timestamps(null: false, type: :utc_datetime_usec)

      # `null` when a pending transaction
      add(:txid, references(:utxotransactions, column: :hash, on_delete: :delete_all, type: :bytea), null: true)

      # `null` when a pending transaction
      add(:block_number, :integer, null: true)

    end

    create(unique_index(:utxoaddresstxs, [:txid, :vout]))
  end
end
