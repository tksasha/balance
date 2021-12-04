# frozen_string_literal: true

RSpec.describe Cashes::DestroyService do
  subject { described_class.new params }

  let(:params) { { id: 16 } }

  describe '#cash' do
    context do
      before { subject.instance_variable_set :@cash, :cash }

      its(:cash) { should eq :cash }
    end

    context do
      before { allow(Cash).to receive(:find).with(16).and_return(:cash) }

      its(:cash) { should eq :cash }
    end
  end

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    context do
      before { allow(cash).to receive(:destroy).and_return(true) }

      its(:call) { should be_success }

      its('call.object') { should eq cash }
    end

    context do
      before { allow(cash).to receive(:destroy).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq cash }
    end
  end
end
