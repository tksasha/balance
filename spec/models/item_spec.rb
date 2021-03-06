# frozen_string_literal: true

RSpec.describe Item, type: :model do
  it { should be_an ActsAsHasFormula }

  it { should be_an ActsAsParanoid }

  it { should validate_presence_of :date }

  it { should validate_presence_of :category_id }

  it { should validate_presence_of :formula }

  it { should validate_presence_of :currency }

  it { should belong_to :category }

  it { should define_enum_for(:currency).with_values(%w[uah usd rub]) }

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
