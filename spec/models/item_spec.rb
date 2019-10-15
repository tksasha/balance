# frozen_string_literal: true

RSpec.describe Item, type: :model do
  it { should be_a ActsAsHasFormula }

  it { should validate_presence_of :date }

  it { should validate_presence_of :category_id }

  it { should validate_presence_of :formula }

  it { should belong_to :category }

  it { should act_as_paranoid }

  describe '.income' do
    let(:sql) { described_class.joins(:category).merge(Category.income).to_sql }

    subject { described_class.income.to_sql }

    it { should eq sql }
  end

  describe '.expense' do
    let(:sql) { described_class.joins(:category).merge(Category.expense).to_sql }

    subject { described_class.expense.to_sql }

    it { should eq sql }
  end
end
