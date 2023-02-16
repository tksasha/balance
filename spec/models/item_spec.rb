# frozen_string_literal: true

RSpec.describe Item do
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

  describe '.for_dashboard' do
    subject { described_class.for_dashboard }

    let(:currencies) { CURRENCIES.keys }

    let(:supercategories) { Category.supercategories.keys }

    let(:items) do
      {
        'uah' => [
          ['uah', 1, 11.11],
          ['uah', 2, 11.11],
          ['uah', 3, 11.11],
          ['uah', 4, 11.11]
        ],
        'usd' => [
          ['usd', 1, 11.11],
          ['usd', 2, 11.11],
          ['usd', 3, 11.11],
          ['usd', 4, 11.11]
        ],
        'eur' => [
          ['eur', 1, 11.11],
          ['eur', 2, 11.11],
          ['eur', 3, 11.11],
          ['eur', 4, 11.11]
        ]
      }
    end

    before do
      travel_to '2023-03-01'

      create(:item, :expense, currency: 'uah', date: '2023-04-01', sum: 22.22)

      create(:item, :income, currency: 'uah', date: '2023-03-01')

      currencies.map do |currency|
        supercategories.map do |supercategory|
          category = create(:category, :expense, currency:, supercategory:, name: Faker::Commerce.product_name)

          create(:item, category:, currency:, sum: 11.11, date: '2023-03-16')
        end
      end
    end

    it { is_expected.to eq items }
  end
end
