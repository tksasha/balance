# frozen_string_literal: true

RSpec.describe ActsAsUpdateAtEndViaWebsocketService do
  let(:described_class) do
    Class.new do
      include ActsAsUpdateAtEndViaWebsocketService

      def currency
        'uah'
      end
    end
  end

  describe '#update_at_end_via_websocket' do
    before { allow(Websockets::UpdateAtEndService).to receive(:call).with('uah') }

    it do
      expect(subject.update_at_end_via_websocket).to eq subject

      expect(Websockets::UpdateAtEndService).to have_received(:call).with('uah')
    end
  end
end
