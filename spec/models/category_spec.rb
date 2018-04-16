require 'rails_helper'

RSpec.describe Category, type: :model do
  it { should validate_presence_of :name }

  it { should validate_uniqueness_of(:name).case_insensitive }

  describe '.group_by_income' do
    before do
      #
      # described_class.visible.expense.pluck(:name, :id) -> [['Їжа', 1], ["Дім. Сім'я", 2]]
      #
      expect(described_class).to receive(:visible) do
        double.tap do |a|
          expect(a).to receive(:expense) do
            double.tap { |b| expect(b).to receive(:pluck).with(:name, :id).and_return([['Їжа', 1], ["Дім. Сім'я", 2]]) }
          end
        end
      end
    end

    before do
      #
      # described_class.visible.income.pluck(:name, :id) -> [['Зарплата', 3]]
      #
      expect(described_class).to receive(:visible) do
        double.tap do |a|
          expect(a).to receive(:income) do
            double.tap { |b| expect(b).to receive(:pluck).with(:name, :id).and_return([['Зарплата', 3]]) }
          end
        end
      end
    end

    subject { described_class.group_by_income }

    it { should eq [['Видатки', [['Їжа', 1], ["Дім. Сім'я", 2]]], ['Надходження', [['Зарплата', 3]]]] }
  end

  describe '#destroy' do
    before { expect(subject).to receive(:update).with(visible: false) }

    it { expect { subject.destroy }.to_not raise_error }
  end

  pending '.visible' do
    before { expect(described_class).to receive(:where).with(visible: true).and_return(:collection) }

    subject { described_class.visible }

    it { should eq :collection }
  end
end
