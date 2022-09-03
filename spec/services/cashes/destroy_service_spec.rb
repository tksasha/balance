# frozen_string_literal: true

RSpec.describe Cashes::DestroyService do
  subject { described_class.new params }

  let(:params) { { id: 16 } }

  it { is_expected.to be_an ActsAsUpdateBalanceViaWebsocketService }

  describe '#cash' do
    context do
      before { subject.instance_variable_set :@cash, :cash }

      its(:cash) { is_expected.to eq :cash }
    end

    context do
      before { allow(Cash).to receive(:find).with(16).and_return(:cash) }

      its(:cash) { is_expected.to eq :cash }
    end
  end

  it { is_expected.to delegate_method(:currency).to(:cash) }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    context do
      before { allow(cash).to receive(:destroy).and_return(true) }

      before { expect(subject).to receive(:update_balance_via_websocket) }

      its(:call) { is_expected.to be_success }

      its('call.object') { is_expected.to eq cash }
    end

    context do
      before { allow(cash).to receive(:destroy).and_return(false) }

      its(:call) { is_expected.to be_failure }

      its('call.object') { is_expected.to eq cash }
    end
  end
end
