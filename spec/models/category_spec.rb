# frozen_string_literal: true

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive }

  describe '.group_by_income' do
    subject { described_class }

    before do
      allow(subject).to \
        receive_message_chain(:visible, :expense, :pluck).with(:name, :id).and_return([['Їжа', 1], ["Дім. Сім'я", 2]])
    end

    before do
      allow(subject).to \
        receive_message_chain(:visible, :income, :pluck).with(:name, :id).and_return([['Зарплата', 3]])
    end

    its :group_by_income do
      should eq [
        ['Видатки', [['Їжа', 1], ["Дім. Сім'я", 2]]],
        ['Надходження', [['Зарплата', 3]]]
      ]
    end
  end

  describe '.visible' do
    let(:sql) { described_class.where(visible: true).to_sql }

    subject { described_class.visible.to_sql }

    it { should eq sql }
  end

  describe '.income' do
    let(:sql) { described_class.where(income: true).to_sql }

    subject { described_class.income.to_sql }

    it { should eq sql }
  end

  describe '.expense' do
    let(:sql) { described_class.where(income: false).to_sql }

    subject { described_class.expense.to_sql }

    it { should eq sql }
  end
end
