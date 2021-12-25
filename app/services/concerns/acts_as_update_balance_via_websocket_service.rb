# frozen_string_literal: true

module ActsAsUpdateBalanceViaWebsocketService
  def update_balance_via_websocket
    ::Websockets::UpdateBalanceService.call(currency)

    self
  end
end
