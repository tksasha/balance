# frozen_string_literal: true

RSpec.describe Items::UpdateService do
  subject { described_class.new params }

  let(:params) do
    acp(
      id: 15,
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

  it { should be_an ActsAsUpdateAtEndViaWebsocketService }

  it { should be_an ActsAsUpdateBalanceViaWebsocketService }

  its(:resource_params) { should eq params.require(:item).permit! }

  describe '#item' do
    context do
      before { subject.instance_variable_set :@item, :item }

      its(:item) { should eq :item }
    end

    context do
      before { allow(Item).to receive(:find).with(15).and_return(:item) }

      its(:item) { should eq :item }
    end
  end

  it { should delegate_method(:currency).to(:item) }

  describe '#call' do
    let(:item) { stub_model Item }

    before { allow(subject).to receive(:item).and_return(item) }

    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    context do
      before { allow(item).to receive(:update).with(:resource_params).and_return(true) }

      before { expect(subject).to receive_message_chain(:update_at_end_via_websocket, :update_balance_via_websocket) }

      its(:call) { should be_success }

      its('call.object') { should eq item }
    end

    context do
      before { allow(item).to receive(:update).with(:resource_params).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq item }
    end
  end
end
