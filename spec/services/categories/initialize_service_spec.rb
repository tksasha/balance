# frozen_string_literal: true

RSpec.describe Categories::InitializeService do
  subject { described_class.new(params) }

  let(:params) { { currency: 'eur' } }

  describe '#category' do
    its(:category) { is_expected.to be_a Category }

    its('category.currency') { is_expected.to eq 'eur' }
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    its(:call) { is_expected.to be_success }

    its('call.object') { is_expected.to eq category }
  end
end
