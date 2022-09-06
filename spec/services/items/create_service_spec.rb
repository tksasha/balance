# frozen_string_literal: true

RSpec.describe Items::CreateService do
  subject { described_class.new params }

  let(:params) do
    acp(
      item: {
        date: nil,
        formula: nil,
        category_id: nil,
        description: nil,
        currency: nil,
        tag_ids: []
      }
    )
  end

  it { is_expected.to be_an ActsAsUpdateAtEndViaWebsocketService }

  it { is_expected.to be_an ActsAsUpdateBalanceViaWebsocketService }

  its(:resource_params) { is_expected.to eq params.require(:item).permit! }

  describe '#item' do
    context 'when @item is set' do
      before { subject.instance_variable_set :@item, :item }

      its(:item) { is_expected.to eq :item }
    end

    context 'when @item is not set' do
      before do
        allow(subject).to receive(:resource_params).and_return(:resource_params)

        allow(Item).to receive(:new).with(:resource_params).and_return(:item)
      end

      its(:item) { is_expected.to eq :item }
    end
  end

  it { is_expected.to delegate_method(:currency).to(:item) }

  describe '#call' do
    let(:item) { stub_model Item }

    before { allow(subject).to receive(:item).and_return(item) }

    context 'when success' do
      before do
        allow(item).to receive(:save).and_return(true)

        allow(subject).to receive_message_chain(:update_at_end_via_websocket, :update_balance_via_websocket)
      end

      its(:call) { is_expected.to be_success }

      its('call.object') { is_expected.to eq item }

      it 'updates via websocket', :aggregate_failures do
        subject.call

        expect(subject).to have_received(:update_at_end_via_websocket)

        expect(subject.update_at_end_via_websocket).to have_received(:update_balance_via_websocket)
      end
    end

    context 'when failure' do
      before { allow(item).to receive(:save).and_return(false) }

      its(:call) { is_expected.to be_failure }

      its('call.object') { is_expected.to eq item }
    end
  end
end
