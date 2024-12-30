defmodule BlockScoutWeb.API.Qitmeer.ApiView do
  def render("message.json", %{message: message}) do
    %{
      "message" => message
    }
  end
end
