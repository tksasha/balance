# frozen_string_literal: true

RSpec.describe Cashes::CreateService do
  subject { described_class.new params }

  let(:params) { acp({ cash: { name: nil, formula: nil, currency: nil, supercategory: nil } }) }

  it { is_expected.to be_an ActsAsUpdateBalanceViaWebsocketService }

  its(:resource_params) { is_expected.to eq params.require(:cash).permit! }

  describe '#cash' do
    context do
      before { subject.instance_variable_set :@cash, :cash }

      its(:cash) { is_expected.to eq :cash }
    end

    context do
      before do
        allow(subject).to receive(:resource_params).and_return(:resource_params)

        allow(Cash).to receive(:new).with(:resource_params).and_return(:cash)
      end

      its(:cash) { is_expected.to eq :cash }
    end
  end

  it { is_expected.to delegate_method(:currency).to(:cash) }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    context do
      before do
        allow(cash).to receive(:save).and_return(true)

        allow(subject).to receive(:update_balance_via_websocket)
      end

      its(:call) { is_expected.to be_success }

      its('call.object') { is_expected.to eq cash }

      it do
        subject.call

        expect(subject).to have_received(:update_balance_via_websocket)
      end
    end

    context do
      before { allow(cash).to receive(:save).and_return(false) }

      its(:call) { is_expected.to be_failure }

      its('call.object') { is_expected.to eq cash }
    end
  end
end
