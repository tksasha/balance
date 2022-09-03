# frozen_string_literal: true

RSpec.describe Cashes::GetCollectionService, type: :cashe do
  subject { described_class.new params }

  let(:params) { double }

  describe '#cashes' do
    before { allow(Cash).to receive(:order).with(:name).and_return(:cashes) }

    its(:cashes) { is_expected.to eq :cashes }
  end

  describe '#call' do
    before { allow(subject).to receive(:cashes).and_return(:cashes) }

    before { allow(CashSearcher).to receive(:call).with(:cashes, params).and_return(:collection) }

    its(:call) { is_expected.to eq :collection }
  end
end
