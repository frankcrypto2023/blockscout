defmodule Indexer.Block.QitmeerCatchup.Fetcher do
  @moduledoc """
  Fetches and indexes block ranges from the block before the latest block to genesis (0) that are missing.
  """

  use Spandex.Decorators

  require Logger

  import Indexer.Block.Fetcher,
    only: [
      qng_fetch_and_import_range: 2
    ]

  import Explorer.Chain.QitmeerBlock, only: [fetch_min_max: 0]
  alias Indexer.Block
  alias Indexer.Block.QitmeerCatchup.TaskSupervisor

  @behaviour Block.Fetcher

  @shutdown_after :timer.minutes(5)

  defstruct block_fetcher: nil,
            memory_monitor: nil

  @impl Block.Fetcher
  def import(_block_fetcher, options) when is_map(options) do
    # db handle
    :ok
  end

  @doc """
  Required named arguments

    * `:json_rpc_named_arguments` - `t:EthereumJSONRPC.json_rpc_named_arguments/0` passed to
        `EthereumJSONRPC.json_rpc/2`.
  """
  def task(state) do
    Logger.metadata(fetcher: :block_catchup)
    stream_fetch_and_import(state)
  end

  defp stream_fetch_and_import(state) do
    TaskSupervisor
    |> Task.Supervisor.async_stream(%{}, &fetch_and_import_range_from_min(state, &1),
      max_concurrency: 1,
      timeout: :infinity,
      shutdown: @shutdown_after
    )
    |> Stream.run()
  end

  defp fetch_and_import_range_from_min(%__MODULE__{block_fetcher: %Block.Fetcher{} = block_fetcher} = state, _) do
    case fetch_min_max() do
      %{min: nil, max: nil} ->
        min = 0
        max_value = 100
        range = min..max_value
        Logger.info(fn -> "Qitmeer Blocks Fetching range #{inspect(range)}" end, fetcher: :block_catchup)
        :timer.tc(fn -> qng_fetch_and_import_range(block_fetcher, range) end)

      %{min: min, max: dbmax} ->
        min = min + 1

        max_value =
          if min + 100 < dbmax do
            min + 100
          else
            dbmax
          end

        range = min..max_value
        Logger.info(fn -> "Qitmeer Blocks Fetching range #{inspect(range)}" end, fetcher: :block_catchup)
        :timer.tc(fn -> qng_fetch_and_import_range(block_fetcher, range) end)
    end

    {:ok, state}
  end
end
