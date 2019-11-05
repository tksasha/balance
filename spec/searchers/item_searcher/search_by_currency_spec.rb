# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  describe '#search_by_currency' do
    subject { described_class.new(relation).search_by_currency(currency) }

    context do
      let(:currency) { '' }

      it { should be_nil }
    end

    context do
      let(:currency) { 'usd' }

      before { expect(relation).to receive(:where).with(currency: 'usd').and_return(:collection) }

      it { should eq :collection }
    end
  end
end
