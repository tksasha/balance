# frozen_string_literal: true

RSpec.describe CategoryWidgetDataSearcher do
  describe '#search' do
    context do
      subject { described_class.search currency: '' }

      it { should eq [] }
    end

    context do
      let :categories do
        [
          ['Food', 13, false],
          ['Drinks', 5, false],
          ['Salary', 17, true],
        ]
      end

      let :collection do
        [
          [
            'Видатки', [
              ['Food', 13],
              ['Drinks', 5],
            ]
          ],
          [
            'Надходження',
            [
              ['Salary', 17],
            ]
          ],
        ]
      end

      before do
        #
        # Category
        #   .where(currency: 'usd')
        #   .order(:income)
        #   .pluck(:name, :id, :income) -> categories
        #
        expect(Category).to receive(:where).with(currency: 'usd') do
          double.tap do |a|
            expect(a).to receive(:order).with(:income) do
              double.tap do |b|
                expect(b).to receive(:pluck).with(:name, :id, :income).and_return(categories)
              end
            end
          end
        end
      end

      subject { described_class.search currency: 'usd' }

      it { should eq collection }
    end
  end

  describe '.search' do
    let(:params) { double }

    after { described_class.search params }

    it do
      #
      # described_class.new(params).search
      #
      expect(described_class).to receive(:new).with(params) do
        double.tap do |a|
          expect(a).to receive(:search)
        end
      end
    end
  end
end
