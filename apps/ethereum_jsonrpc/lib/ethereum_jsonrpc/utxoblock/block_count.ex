defmodule EthereumJSONRPC.QitmeerBlock.BlockCount do
  @moduledoc false
  def request(json_rpc_named_arguments) do
    EthereumJSONRPC.request(%{id: 1, method: "qng_getBlockCount", params: []})
    |> EthereumJSONRPC.json_rpc(json_rpc_named_arguments)
  end

  @moduledoc false
  def number_from_result({:ok, count}) do
    {:ok, count}
  end
end
