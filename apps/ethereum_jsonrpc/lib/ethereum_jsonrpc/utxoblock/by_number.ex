defmodule EthereumJSONRPC.UTXOBlock.ByNumber do
  @moduledoc """
  Block format as returned by [`qng_getBlockByOrder`]
  """

  import EthereumJSONRPC, only: [integer_to_quantity: 1]

  def request(%{id: id, number: number}) do
    EthereumJSONRPC.request(%{id: id, method: "qng_getBlockByOrder", params: [number, true]})
  end
end
