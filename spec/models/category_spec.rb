# frozen_string_literal: true

RSpec.describe Category do
  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(uah: 0, usd: 1, eur: 3) }

  it do
    expect(subject.defined_enums)
      .to include 'supercategory' => {
        'common' => 1, 'children' => 2, 'business' => 3, 'charity' => 4, 'currency' => 5
      }
  end

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
