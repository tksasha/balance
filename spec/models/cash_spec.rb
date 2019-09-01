# frozen_string_literal: true

RSpec.describe Cash, type: :model do
  subject { Cash.new sum: 0 }

  it { should act_as_paranoid }

  it { should be_a ActsAsHasFormula }

  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive }

  it { should validate_presence_of :formula }

  describe '.at_end' do
    before do
      #
      # Item.income.sum(:sum) -> 10
      #
      expect(Item).to receive(:income) do
        double.tap { |a| expect(a).to receive(:sum).with(:sum) { 10 } }
      end
    end

    before do
      #
      # Item.expense.sum(:sum) -> 6.5
      #
      expect(Item).to receive(:expense) do
        double.tap { |b| expect(b).to receive(:sum).with(:sum) { 6.5 } }
      end
    end

    subject { described_class.at_end }

    it { should eq 3.5 }
  end

  describe '.balance' do
    before { expect(described_class).to receive(:sum).with(:sum).and_return(24) }

    before { expect(described_class).to receive(:at_end).and_return(19.8) }

    subject { described_class.balance }

    it { should eq 4.2 }
  end
end
