# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  describe '#search_by_month' do
    let(:month) { Month.new(2021, 3) }

    let(:dates) { Date.new(2021, 3, 1)..Date.new(2021, 3, 31) }

    let(:collection) { double }

    subject { described_class.new(relation).search_by_month(month) }

    before { allow(relation).to receive(:where).with(date: dates).and_return(collection) }

    it { should eq collection }
  end
end
