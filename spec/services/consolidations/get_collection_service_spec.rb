# frozen_string_literal: true

RSpec.describe Consolidations::GetCollectionService do
  subject { described_class.new params }

  let(:params) { double }

  describe '#scope', skip: 'private method' do
    before { allow(Consolidation).to receive(:includes).with(:category).and_return(:scope) }

    its(:scope) { is_expected.to eq :scope }
  end

  describe '#call' do
    before do
      allow(subject).to receive(:scope).and_return(:scope)

      allow(ConsolidationSearcher).to receive(:search).with(:scope, params).and_return(:collection)
    end

    its(:call) { is_expected.to eq :collection }
  end
end
