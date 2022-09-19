# frozen_string_literal: true

RSpec.describe Item, type: :model do
  it { is_expected.to be_an ActsAsHasFormula }

  it { is_expected.to be_an ActsAsParanoid }

  it { is_expected.to belong_to(:category).required }

  it { is_expected.to validate_presence_of :date }

  it { is_expected.to validate_presence_of :formula }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(uah: 0, usd: 1, eur: 3) }

  describe '.income' do
    subject { described_class.income.to_sql }

    let(:sql) { described_class.joins(:category).merge(Category.income).to_sql }

    it { is_expected.to eq sql }
  end

  describe '.expense' do
    subject { described_class.expense.to_sql }

    let(:sql) { described_class.joins(:category).merge(Category.expense).to_sql }

    it { is_expected.to eq sql }
  end
end
