# frozen_string_literal: true

RSpec.describe Cashes::InitializeService do
  subject { described_class.new(params) }

  let(:params) { { currency: 'usd' } }

  its(:cash) { should be_a Cash }

  its('cash.currency') { should eq 'usd' }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    its(:call) { should be_success }

    its('call.object') { should eq cash }
  end
end
