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
    before { expect(Websockets::UpdateAtEndService).to receive(:call).with('uah') }

    its(:update_at_end_via_websocket) { should eq subject }
  end
end
