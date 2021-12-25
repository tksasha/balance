# frozen_string_literal: true

RSpec.describe ActsAsUpdateBalanceViaWebsocketService do
  let(:described_class) do
    Class.new do
      include ActsAsUpdateBalanceViaWebsocketService

      def currency
        'uah'
      end
    end
  end

  describe '#update_balance_via_websocket' do
    before { expect(Websockets::UpdateBalanceService).to receive(:call).with('uah') }

    its(:update_balance_via_websocket) { should eq subject }
  end
end
