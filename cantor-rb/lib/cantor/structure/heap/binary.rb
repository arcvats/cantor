# frozen_string_literal: true

module Structure
  module Heap
    # Defines a Binary heap (default is min heap)
    class BinaryHeap
      attr_reader :size

      def initialize(comparable = nil)
        @elements = []
        @comparable = comparable || ->(a, b) { a < b }
        @size = 0
      end

      def empty?
        @size.zero?
      end

      def peek
        return nil if empty?

        @elements[0]
      end

      def swap(idx, idy)
        @elements[idx], @elements[idy] = @elements[idy], @elements[idx]
      end

      def insert(element)
        @size += 1
        @elements << element
        idx = @size - 1
        parent_idx = idx / 2

        while idx.positive? && @comparable.call(@elements[idx], @elements[parent_idx])
          swap(idx, parent_idx)
          idx = parent_idx
          parent_idx = idx / 2
        end
      end

      def delete
        return nil if empty?

        element = @elements[0]
        swap(0, @size - 1)
        @size -= 1
        @elements = @elements[0...@size]
        idx = 0

        while idx < @size / 2
          left = idx * 2 + 1
          right = idx * 2 + 2

          right_with_left = @comparable.call(@elements[right], @elements[left])
          right_with_idx = @comparable.call(@elements[right], @elements[idx])
          left_with_idx = @comparable.call(@elements[left], @elements[idx])
          if right < @size && right_with_left && right_with_idx
            swap(idx, right)
            idx = right
          elsif left_with_idx
            swap(idx, left)
            idx = left
          else
            break
          end
        end
        element
      end

      def sort
        array = []
        array << delete until @size.zero?
        array
      end
    end
  end
end
