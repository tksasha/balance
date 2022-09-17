# frozen_string_literal: true

RSpec.describe Category, type: :model do
  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(%w[uah usd rub eur]) }

  it { is_expected.to define_enum_for(:supercategory).with_values(first: 1, second: 2, third: 3) }

  describe '.visible' do
    subject { described_class.visible.to_sql }

    let(:sql) { described_class.where(visible: true).to_sql }

    it { is_expected.to eq sql }
  end

  describe '.income' do
    subject { described_class.income.to_sql }

    let(:sql) { described_class.where(income: true).to_sql }

    it { is_expected.to eq sql }
  end

  describe '.expense' do
    subject { described_class.expense.to_sql }

    let(:sql) { described_class.where(income: false).to_sql }

    it { is_expected.to eq sql }
  end
end
