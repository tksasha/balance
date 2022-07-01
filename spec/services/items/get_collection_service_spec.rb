# frozen_string_literal: true

RSpec.describe Items::GetCollectionService do
  subject { described_class.new params }

  let(:params) { double }

  describe '#items' do
    before do
      #
      # Item
      # .order(date: :desc)
      # .includes(:category) -> :items
      #
      allow(Item).to receive(:order).with(date: :desc) do
        double.tap do |a|
          allow(a).to receive(:includes).with(:category).and_return(:items)
        end
      end
    end

    its(:items) { should eq :items }
  end

  describe '#call' do
    before { allow(subject).to receive(:items).and_return(:items) }

    before { allow(ItemSearcher).to receive(:call).with(:items, params).and_return(:collection) }

    its(:call) { should eq :collection }
  end
end
