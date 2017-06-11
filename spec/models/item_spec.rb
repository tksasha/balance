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
    before do
      #
      # described_class.
      #   joins('INNER JOIN categories ON categories.id=items.category_id').
      #   where({ categories: { income: true } }) -> :collection
      #
      expect(described_class).to receive(:joins).with('INNER JOIN categories ON categories.id=items.category_id') do
        double.tap { |a| expect(a).to receive(:where).with({ categories: { income: true } }).and_return(:collection) }
      end
    end

    subject { described_class.income }

    it { should eq :collection }
  end

  describe '.expense' do
    before do
      #
      # described_class.
      #   joins('INNER JOIN categories ON categories.id=items.category_id').
      #   where({ categories: { income: false } }) -> :collection
      #
      expect(described_class).to receive(:joins).with('INNER JOIN categories ON categories.id=items.category_id') do
        double.tap { |a| expect(a).to receive(:where).with({ categories: { income: false } }).and_return(:collection) }
      end
    end

    subject { described_class.expense }

    it { should eq :collection }
  end
end
