require 'rails_helper'

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive }

  describe '.visible' do
    before { expect(Category).to receive(:where).with(visible: true) }

    it { expect { Category.visible }.to_not raise_error }
  end

  describe '.group_by_income' do
    before do
      #
      # Category.visible.expense.pluck(:name, :id) -> [['Їжа', 1], ["Дім. Сім'я", 2]]
      #
      expect(Category).to receive(:visible) do
        double.tap do |a|
          expect(a).to receive(:expense) do
            double.tap { |b| expect(b).to receive(:pluck).with(:name, :id).and_return([['Їжа', 1], ["Дім. Сім'я", 2]]) }
          end
        end
      end
    end

    before do
      #
      # Category.visible.income.pluck(:name, :id) -> [['Зарплата', 3]]
      #
      expect(Category).to receive(:visible) do
        double.tap do |a|
          expect(a).to receive(:income) do
            double.tap { |b| expect(b).to receive(:pluck).with(:name, :id).and_return([['Зарплата', 3]]) }
          end
        end
      end
    end

    it do
      expect(Category.group_by_income).
        to eq [['Видатки', [['Їжа', 1], ["Дім. Сім'я", 2]]], ['Надходження', [['Зарплата', 3]]]]
    end
  end

  describe '#destroy' do
    before { expect(subject).to receive(:update).with(visible: false) }

    it { expect { subject.destroy }.to_not raise_error }
  end
end
