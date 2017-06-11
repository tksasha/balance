require 'rails_helper'

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive }

  it { should act_as_paranoid }

  describe '.group_by_income' do
    before do
      #
      # described_class.expense.pluck(:name, :id) -> [['Їжа', 1], ["Дім. Сім'я", 2]]
      #
      expect(described_class).to receive(:expense) do
        double.tap { |a| expect(a).to receive(:pluck).with(:name, :id).and_return([['Їжа', 1], ["Дім. Сім'я", 2]]) }
      end
    end

    before do
      #
      # described_class.income.pluck(:name, :id) -> [['Зарплата', 3]]
      #
      expect(described_class).to receive(:income) do
        double.tap { |a| expect(a).to receive(:pluck).with(:name, :id).and_return([['Зарплата', 3]]) }
      end
    end

    subject { described_class.group_by_income }

    it { should eq [['Видатки', [['Їжа', 1], ["Дім. Сім'я", 2]]], ['Надходження', [['Зарплата', 3]]]] }
  end
end
