# frozen_string_literal: true

RSpec.describe Items::DestroyService do
  subject { described_class.new params }

  let(:params) { { id: 25 } }

  it { is_expected.to be_an ActsAsUpdateAtEndViaWebsocketService }

  it { is_expected.to be_an ActsAsUpdateBalanceViaWebsocketService }

  describe '#item' do
    context do
      before { subject.instance_variable_set :@item, :item }

      its(:item) { is_expected.to eq :item }
    end

    context do
      before { allow(Item).to receive(:find).with(25).and_return(:item) }

      its(:item) { is_expected.to eq :item }
    end
  end

  it { is_expected.to delegate_method(:currency).to(:item) }

  describe '#call' do
    let(:item) { stub_model Item }

    before { allow(subject).to receive(:item).and_return(item) }

    before { allow(item).to receive(:destroy).and_return(true) }

    before { expect(subject).to receive_message_chain(:update_at_end_via_websocket, :update_balance_via_websocket) }

    its(:call) { is_expected.to be_success }

    its('call.object') { is_expected.to eq item }
  end
end
