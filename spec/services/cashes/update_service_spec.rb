# frozen_string_literal: true

RSpec.describe Cashes::UpdateService, type: :service do
  subject { described_class.new params }

  let(:params) { acp({ id: 15, cash: { name: nil, formula: nil, currency: nil } }) }

  its(:resource_params) { should eq params.require(:cash).permit! }

  describe '#cash' do
    context do
      before { subject.instance_variable_set :@cash, :cash }

      its(:cash) { should eq :cash }
    end

    context do
      before { allow(Cash).to receive(:find).with(15).and_return(:cash) }

      its(:cash) { should eq :cash }
    end
  end

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    context do
      before { allow(cash).to receive(:update).with(:resource_params).and_return(true) }

      its(:call) { should be_success }

      its('call.object') { should eq cash }
    end

    context do
      before { allow(cash).to receive(:update).with(:resource_params).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq cash }
    end
  end
end
