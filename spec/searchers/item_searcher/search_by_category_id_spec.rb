# frozen_string_literal: true

RSpec.describe ItemSearcher do
  describe '#search_by_category_id' do
    subject { described_class.new(relation).search_by_category_id(category_id) }

    let(:relation) { double }

    context do
      let(:category_id) { '' }

      it { should be_nil }
    end

    context do
      let(:category_id) { 9053 }

      let(:collection) { double }

      before { allow(relation).to receive(:where).with(category_id: 9053).and_return(collection) }

      it { should eq collection }
    end
  end
end
