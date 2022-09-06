# frozen_string_literal: true

RSpec.describe Categories::GetCollectionService do
  subject { described_class.new params }

  let(:params) { double }

  describe '#categories' do
    before { allow(Category).to receive(:order).with(:income).and_return(:categories) }

    its(:categories) { is_expected.to eq :categories }
  end

  describe '#call' do
    before do
      allow(subject).to receive(:categories).and_return(:categories)

      allow(CategorySearcher).to receive(:call).with(:categories, params).and_return(:collection)
    end

    its(:call) { is_expected.to eq :collection }
  end
end
