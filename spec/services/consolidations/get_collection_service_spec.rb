# frozen_string_literal: true

RSpec.describe Consolidations::GetCollectionService do
  subject { described_class.new params }

  let(:params) { double }

  describe '#scope' do
    before { allow(Consolidation).to receive(:includes).with(:category).and_return(:scope) }

    its(:scope) { is_expected.to eq :scope }
  end

  describe '#call' do
    before { allow(subject).to receive(:scope).and_return(:scope) }

    before { allow(ConsolidationSearcher).to receive(:search).with(:scope, params).and_return(:collection) }

    its(:call) { is_expected.to eq :collection }
  end
end
