# frozen_string_literal: true

RSpec.describe ItemSearcher do
  let(:relation) { double }

  subject { described_class.new relation }

  describe '#results' do
    before { expect(relation).to receive(:includes).with(:category).and_return(:results) }

    its(:results) { should eq :results }
  end
end
