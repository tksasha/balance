# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  let(:date_range) { double }

  before { allow(subject).to receive(:date_range).and_return(date_range) }

  subject { described_class.new relation, category: slug }

  describe '#search_by_category' do
    context do
      let(:slug) { '' }

      before do
        expect(relation).to \
          receive_message_chain(:includes, :where)
          .with(:category)
          .with(date: date_range)
          .and_return(:collection)
      end

      its(:search) { should eq :collection }
    end

    context do
      let(:slug) { 'drinks' }

      before do
        expect(relation).to \
          receive_message_chain(:includes, :where, :where)
          .with(:category)
          .with(categories: { slug: 'drinks' })
          .with(date: date_range)
          .and_return(:collection)
      end

      its(:search) { should eq :collection }
    end
  end
end
