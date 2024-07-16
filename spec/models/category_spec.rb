# frozen_string_literal: true

# == Schema Information
#
# Table name: categories
#
#  id            :integer          not null, primary key
#  currency      :integer          default("uah")
#  deleted_at    :datetime
#  income        :boolean          default(FALSE)
#  name          :string
#  supercategory :integer          default("common"), not null
#  visible       :boolean          default(TRUE)
#
# Indexes
#
#  index_categories_on_name_and_currency  (name,currency) UNIQUE
#
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

  describe '#destroy' do
    subject { described_class.new }

    before { allow(subject).to receive(:touch).with(:deleted_at) }

    it { expect { subject.destroy }.to have_received(:touch).with(:deleted_at) }
  end

  describe 'default scope' do
    let!(:category_one) { create(:category, deleted_at: nil) }

    let!(:category_two) { create(:category, deleted_at: Time.zone.now) }

    let(:collection) { described_class.all }

    it { expect(collection).to include(category_one) }

    it { expect(collection).not_to include(category_two) }
  end
end
