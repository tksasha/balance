# frozen_string_literal: true

RSpec.describe Categories::InitializeService do
  subject { described_class.new(params) }

  let(:params) { { currency: 'eur' } }

  describe '#category' do
    its(:category) { should be_a Category }

    its('category.currency') { should eq 'eur' }
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    its(:call) { should be_success }

    its('call.object') { should eq category }
  end
end
