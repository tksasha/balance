# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  describe '#search_by_category' do
    subject { described_class.new(relation).search_by_category(category) }

    context do
      let(:category) { '' }

      it { should be_nil }
    end

    context do
      let(:category) { 'drinks' }

      before do
        #
        # relation
        #   .joins(:category)
        #   .where(categories: { slug: 'drinks' }) -> :collection
        #
        expect(relation).to receive(:joins).with(:category) do
          double.tap do |a|
            expect(a).to receive(:where).with(categories: { slug: 'drinks' }).and_return(:collection)
          end
        end
      end

      it { should eq :collection }
    end
  end
end
