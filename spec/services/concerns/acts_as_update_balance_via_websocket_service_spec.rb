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
    before { allow(Websockets::UpdateBalanceService).to receive(:call).with('uah') }

    it do
      expect(subject.update_balance_via_websocket).to eq subject

      expect(Websockets::UpdateBalanceService).to have_received(:call).with('uah')
    end
  end
end
