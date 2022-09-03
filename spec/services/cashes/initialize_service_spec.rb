# frozen_string_literal: true

RSpec.describe Cashes::InitializeService do
  subject { described_class.new(params) }

  let(:params) { { currency: 'usd' } }

  its(:cash) { is_expected.to be_a Cash }

  its('cash.currency') { is_expected.to eq 'usd' }

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    its(:call) { is_expected.to be_success }

    its('call.object') { is_expected.to eq cash }
  end
end
