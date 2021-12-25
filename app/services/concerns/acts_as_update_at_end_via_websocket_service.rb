# frozen_string_literal: true

module ActsAsUpdateAtEndViaWebsocketService
  def update_at_end_via_websocket
    ::Websockets::UpdateAtEndService.call(currency)

    self
  end
end
