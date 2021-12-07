# frozen_string_literal: true

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { should validate_presence_of :currency }

  it { should define_enum_for(:currency).with_values(%w[uah usd rub]) }

  describe '.visible' do
    subject { described_class.visible.to_sql }

    let(:sql) { described_class.where(visible: true).to_sql }

    it { should eq sql }
  end

  describe '.income' do
    subject { described_class.income.to_sql }

    let(:sql) { described_class.where(income: true).to_sql }

    it { should eq sql }
  end

  describe '.expense' do
    subject { described_class.expense.to_sql }

    let(:sql) { described_class.where(income: false).to_sql }

    it { should eq sql }
  end
end
