# frozen_string_literal: true

RSpec.describe Cashes::CreateService do
  subject { described_class.new params }

  let(:params) { acp({ cash: { name: nil, formula: nil, currency: nil } }) }

  it { should be_an ActsAsUpdateBalanceViaWebsocketService }

  its(:resource_params) { should eq params.require(:cash).permit! }

  describe '#cash' do
    context do
      before { subject.instance_variable_set :@cash, :cash }

      its(:cash) { should eq :cash }
    end

    context do
      before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

      before { allow(Cash).to receive(:new).with(:resource_params).and_return(:cash) }

      its(:cash) { should eq :cash }
    end
  end

  it { should delegate_method(:currency).to(:cash) }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    context do
      before { allow(cash).to receive(:save).and_return(true) }

      before { expect(subject).to receive(:update_balance_via_websocket) }

      its(:call) { should be_success }

      its('call.object') { should eq cash }
    end

    context do
      before { allow(cash).to receive(:save).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq cash }
    end
  end
end
