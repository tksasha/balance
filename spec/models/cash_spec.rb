# frozen_string_literal: true

# == Schema Information
#
# Table name: cashes
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  deleted_at    :time
#  favorite      :boolean          default(FALSE)
#  formula       :string
#  name          :string
#  sum           :decimal(10, 2)
#  supercategory :integer          default("cash"), not null
#
# Indexes
#
#  index_cashes_on_name_and_currency  (name,currency) UNIQUE
#
RSpec.describe Cash do
  it { is_expected.to be_an ActsAsHasFormula }

  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :formula }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(uah: 0, usd: 1, eur: 3) }

  it { expect(subject.defined_enums).to include 'supercategory' => { 'cash' => 1, 'bonds' => 2, 'deposits' => 3 } }

  it { is_expected.to be_versioned }

  describe '.for_dashboard' do
    subject { described_class.for_dashboard }

    let(:currencies) { CURRENCIES.keys }

    let(:supercategories) { described_class.supercategories.keys }

    let(:cashes) do
      {
        'uah' => [
          ['uah', 'cash', 11.11],
          ['uah', 'bonds', 11.11],
          ['uah', 'deposits', 11.11]
        ],
        'usd' => [
          ['usd', 'cash', 11.11],
          ['usd', 'bonds', 11.11],
          ['usd', 'deposits', 11.11]
        ],
        'eur' => [
          ['eur', 'cash', 11.11],
          ['eur', 'bonds', 11.11],
          ['eur', 'deposits', 11.11]
        ]
      }
    end

    before do
      currencies.map do |currency|
        supercategories.map do |supercategory|
          create(:cash, currency:, supercategory:, sum: 11.11)
        end
      end
    end

    it { is_expected.to eq cashes }
  end
end
