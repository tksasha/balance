# frozen_string_literal: true

RSpec.describe Category do
  it { is_expected.to validate_presence_of :name }

  it { is_expected.to validate_uniqueness_of(:name).case_insensitive.scoped_to(:currency) }

  it { is_expected.to validate_presence_of :currency }

  it { is_expected.to define_enum_for(:currency).with_values(uah: 0, usd: 1, eur: 3) }

  it { expect(subject.defined_enums).to include 'supercategory' => { 'one' => 1, 'two' => 2, 'three' => 3 } }

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

  describe '.supercategories' do
    subject { described_class.supercategories }

    it do
      I18n.with_locale(:en) do
        expect(subject).to eq 'First' => 'one', 'Second' => 'two', 'Third' => 'three'
      end
    end

    it do
      I18n.with_locale(:ua) do
        expect(subject).to eq 'Перша' => 'one', 'Друга' => 'two', 'Третя' => 'three'
      end
    end
  end
end
