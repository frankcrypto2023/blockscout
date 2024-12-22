defmodule Explorer.Repo.Migrations.CreateUTXOAddressBalance do
  use Ecto.Migration

  def change do
    create table(:utxoaddressbalance, primary_key: false) do
      add(:address, :string, null: true)
      add(:available, :bigint, null: false)
      add(:unavailable, :bigint, null: false)
      add(:spent, :bigint, null: false)
      timestamps(null: false, type: :utc_datetime_usec)

    end

    create(unique_index(:utxoaddressbalance, [:address]))
  end
end
