require 'rails_helper'

RSpec.describe Item, type: :model do
  it { should be_a ActsAsHasFormula }

  it { should validate_presence_of :date }

  it { should validate_presence_of :category_id }

  it { should validate_presence_of :formula }

  it { should belong_to :category }

  it { should act_as_paranoid }

  describe '.search' do
    let(:date_range) { DateRange.new Date.today }

    context do
      before do
        #
        # Item.includes(:category).where(date: date_range).order('date DESC')
        #
        expect(Item).to receive(:includes).with(:category) do
          double.tap do |a|
            expect(a).to receive(:where).with(date: date_range) do
              double.tap { |b| expect(b).to receive(:order).with('date DESC') }
            end
          end
        end
      end

      it { expect { Item.search date_range }.to_not raise_error }
    end

    context do
      before do
        #
        # Item.where('categories.slug' => 'food').includes(:category).where(date: date_range).order('date DESC')
        #
        expect(Item).to receive(:where).with('categories.slug' => 'food') do
          double.tap do |a|
            expect(a).to receive(:includes).with(:category) do
              double.tap do |b|
                expect(b).to receive(:where).with(date: date_range) do
                  double.tap { |c| expect(c).to receive(:order).with('date DESC') }
                end
              end
            end
          end
        end
      end

      it { expect { Item.search date_range, 'food' }.to_not raise_error }
    end
  end

  describe '.income' do
    let(:sql) { described_class.joins(:category).merge(Category.income).to_sql }

    subject { described_class.income.to_sql }

    it { should eq sql }
  end

  describe '.expense' do
    let(:sql) { described_class.joins(:category).merge(Category.expense).to_sql }

    subject { described_class.expense.to_sql }

    it { should eq sql }
  end
end
